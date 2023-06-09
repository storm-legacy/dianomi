import React, { useState } from 'react';

export const VideoComment = () => {
  const [comment, setComment] = useState('');

  const handleCommentChange = (event: any) => {
    setComment(event.target.value);
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();

    setComment('');
  };
  return (
    <div className="row align-items-start" style={{ maxHeight: '50vh' }}>
      <div className="col-3 ">
        <div className="panel panel-default">
          <div className="panel-heading">Dodaj komentarz</div>
          <div className="panel-body">
            <form onSubmit={handleSubmit}>
              <div className="form-group">
                <textarea
                  className="form-control"
                  value={comment}
                  rows={3}
                  maxLength={250}
                  onChange={handleCommentChange}
                />
              </div>
              <button type="submit" className="btn btn-primary">
                Dodaj komentarz
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};
