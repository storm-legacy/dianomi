import React, { useState } from 'react';
import VideoService, { CategoriesName } from '../../../services/video.service';

function CategoriAdd() {
  const [categories, setCategories] = useState('');
  const handleSubmit = async (event: any) => {
    event.preventDefault();
    console.log(categories);
    const data: CategoriesName = {
      name: categories,
    };
    const { request } = VideoService.sendCategori(data);
    request
      .then((res) => {
        window.location.reload();
      })
      .catch((err) => {
        console.log(err.message);
      });
  };
  return (
    <>
      <div className="text-center">
        <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
          <form className="row" onSubmit={handleSubmit}>
            {' '}
            <label>
              <h3>Add categories </h3>
              <input
                className="form-control"
                type="text"
                value={categories}
                onChange={(event) => setCategories(event.target.value)}
              />
            </label>
            <p></p>
            <button type="submit" className="btn btn-primary ">
              Send
            </button>
          </form>
        </div>
      </div>
    </>
  );
}

export default CategoriAdd;
