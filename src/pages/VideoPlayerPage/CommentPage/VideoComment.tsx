import React, { useContext, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import videoService, { CommentData } from '../../../services/video.service';
import { AuthContext } from '../../../context/AuthContext';

export const VideoComment = () => {
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  const [comment, setComment] = useState('');
  const [comments, setComments] = useState([{ email: '', comment: '', updated_at: '' }]);
  const { user } = useContext(AuthContext);

  const handleCommentChange = (event: any) => {
    setComment(event.target.value);
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();
    const data: CommentData = {
      email: user?.email,
      video_id: VideoIdInt,
      comment: comment,
    };
    if (comment != null) {
      const { request } = videoService.sendComment(data);
      request
        .then((res) => {
          setComment('');
        })
        .catch((err) => {
          console.error(err);
        });
    }
  };

  useEffect(() => {
    const { request } = videoService.takeCommentVideoId(VideoIdInt);
    request
      .then((res) => {
        setComments(res.data);
        console.log(comments);
      })
      .catch((err) => {
        console.error(err);
      });
  }, [handleSubmit]);
  return (
    <div className="col-4 position-fixed top-0 end-0 mt-5 me-5 ">
      <div className="panel panel-default">
        <div className="panel-heading">Dodaj komentarz</div>
        <div className="panel-body">
          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <textarea
                style={{ resize: 'none' }}
                className="form-control"
                value={comment}
                rows={3}
                maxLength={250}
                onChange={handleCommentChange}
              />
            </div>
            <button type="submit" className="btn btn-primary mb-3">
              Dodaj komentarz
            </button>
          </form>
        </div>
      </div>
      <div className="overflow-auto" style={{ height: '75vh' }}>
        {comments.map((item, index) => (
          <div className="comment card mb-3 me-3" key={index}>
            <div className="comment-header card-header">
              <h4 className="comment-author card-title">{item.email}</h4>
              <span className="comment-date">{item.updated_at.slice(0, 10)}</span>
            </div>
            <div className="comment-body card-body">
              <p className="comment-text card-text">{item.comment}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
