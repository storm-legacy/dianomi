import React, { useContext, useEffect, useState, FormEvent } from 'react';
import { useParams } from 'react-router-dom';
import videoService, { CommentData } from '../../../services/video.service';
import { AuthContext } from '../../../context/AuthContext';
import { FaFlag } from 'react-icons/fa';
import { Confirm, Notify } from 'notiflix';

export const VideoComment = () => {
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  const [comment, setComment] = useState('');
  const [comments, setComments] = useState([{ id: '', email: '', comment: '', updated_at: '' }]);
  const { user } = useContext(AuthContext);
  const [reset, setReset] = useState(false);

  const handleCommentChange = (event: any) => {
    setComment(event.target.value);
  };

  useEffect(() => {
    setReset(false);
    const { request } = videoService.takeCommentVideoId(VideoIdInt);
    request
      .then((res) => {
        setComments(res.data);
      })
      .catch((err) => {
        console.error(err);
      });
  }, [reset]);

  const handleSubmit = (event: any) => {
    event.preventDefault();
    const data: CommentData = {
      email: user?.email,
      video_id: VideoIdInt,
      comment: comment,
    };

    setReset(true);
    const { request } = videoService.sendComment(data);
    request
      .then((res) => {
        setComment('');
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const handleReport = (commentId: number) => {
    Confirm.prompt(
      'Report comment',
      'What is the reason?',
      'Message is inappropriate',
      'Confirm',
      'Cancel',
      (reportMessage) => {
        if (reportMessage.length < 8) {
          Notify.warning('Report message was too short.');
          return;
        }
        const { request } = videoService.reportComment(commentId, reportMessage);
        request
          .then(() => {
            Notify.success('Comment successfully reported!');
          })
          .catch(() => {
            Notify.failure('Error occured while trying to report comment');
          });
      },
      () => {
        Notify.warning('Report was canceled.');
      },
    );
  };

  return (
    <div className="col-3 position-fixed top-0 end-0 mt-5 me-5 ">
      <div className="panel panel-default">
        <div className="panel-heading">Add comment</div>
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
              Add
            </button>
          </form>
        </div>
      </div>
      <div className="overflow-auto" style={{ height: '75vh' }}>
        {comments.map((item, index) => (
          <div className="comment card mb-3 me-3" key={index}>
            <div className="comment-header card-header position-relative">
              <button
                className="btn btn-danger position-absolute top-50 start-100 translate-middle"
                onClick={() => handleReport(Number(item.id))}
              >
                <FaFlag />
              </button>
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
