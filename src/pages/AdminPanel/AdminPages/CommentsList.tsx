import React, { useEffect, useState } from 'react';
import { FaTrash } from 'react-icons/fa';
import videoService from '../../../services/video.service';
import Paginate from '../../../components/Paginate';
import { CommentReport, Comment } from '../../../types/comment.type';
import { Notify } from 'notiflix';



export const CommentsList = () => {
  const [comments, setComments] = useState<Comment[]>([]);

  const [currentPage, setCurrentPage] = useState(1);
  const [commentsPerPage] = useState(4);

  const indexOfLastVideo = currentPage * commentsPerPage;
  const indexOfFirstVideo = indexOfLastVideo - commentsPerPage;
  const currentComments = comments.slice(indexOfFirstVideo, indexOfLastVideo);

  const refreshComments = () => {
    const { request } = videoService.takeAllComment();
    request
      .then((res) => {
        setComments(res.data);
      })
      .catch((err) => {
        Notify.failure("Comments could not be downloaded!");
      });
  }

  useEffect(() => {
    refreshComments();
  }, []);

  const closeReport = (reportId: number) => {
    const { request } = videoService.closeCommentReport(reportId);
    request
      .then((res) => {
        refreshComments();
        Notify.success("Report was closed!");
      })
      .catch((err) => {
        Notify.failure("Could not close the report");
      });
  }

  const DeleteComment = (id: number) => {
    const { request } = videoService.deleteComment(id);
    request.then(() => {
      Notify.success("Comment was successfully removed!");
      refreshComments();
    }).catch(() => {
      Notify.failure("Comment could not be removed!");
    })
  };

  return (
    <>
      {currentComments.map((item, index) => (
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
          {
            item.reports && (
              <div className="p-2 border border-solid bg-warning mouse-click">
                {item.reports && item.reports.map((report, i) => (
                  <div style={{ cursor: "pointer" }} onClick={() => closeReport(report.ID)}>
                    <span key={i}>{report.Message}</span>
                  </div>
                ))}
              </div>
            )
          }
        </div>
      ))}
      <div className="d-flex justify-content-center p-4">
        <Paginate postsPerPage={6} totalPosts={comments.length} paginate={setCurrentPage} />
      </div>
    </>
  );
};
