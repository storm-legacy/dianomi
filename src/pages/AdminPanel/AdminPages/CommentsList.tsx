import React, { useEffect, useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import videoService from '../../../services/video.service';

export const CommentsList = () => {
  const [comments, setComments] = useState([{ id: 0, name: '', email: '', comment: '', updated_at: '' }]);
  const { request } = videoService.takeAllComment();

  useEffect(() => {
    request
      .then((res) => {
        setComments(res.data);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);
  const DeleteComment = (id: number) => {
    console.log(id);
    const { request } = videoService.deleteComment(id);
    request
      .then(() => window.location.reload())
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <>
      {comments.map((item, index) => (
        <div className="comment card mb-3 me-3" key={index}>
          <div className="comment-header card-header">
            <h4 className="comment-author card-title">
              {item.email} | {item.name}
            </h4>
            <span className="comment-date">{item.updated_at.slice(0, 10)}</span>
            <span className="comment-date">
              {' '}
              <button className="btn btn-danger mx-2" onClick={() => DeleteComment(item.id)}>
                <FaTrash />
              </button>{' '}
            </span>
          </div>
          <div className="comment-body card-body">
            <p className="comment-text card-text">{item.comment}</p>
          </div>
        </div>
      ))}
    </>
  );
};
