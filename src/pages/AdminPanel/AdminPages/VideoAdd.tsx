import React, { useState } from 'react';

export const VideoAdd = () => {
  const [videoFile, setVideoFile] = useState('');

  const handleSubmit = (event: any) => {
    event.preventDefault();

    const videoData = {
      videoFile: videoFile,
    };
    console.log({ videoData });
  };
  return (
    <>
      <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
        <h3>Dodaj Wideo</h3>
        <p></p>
        <form onSubmit={handleSubmit} className="row">
          <label>
            <p>Nazwa</p>
            <input className="form-control" type="text" />
          </label>
          <label>
            <p>Opis</p>
            <input className="form-control" type="text" />
          </label>
          <label>
            <p>Plik</p>
            <input className="form-control" type="file" aria-describedby="basic-addon1" value={videoFile} />
          </label>
          <label>
            <p>Autor</p>
            <input className="form-control" type="text" />
          </label>
          <label>
            <p>Kategoria</p>
            <input className="form-control" type="text" />
          </label>
          <label>
            <p>tag</p>
            <input className="form-control" type="text" />
          </label>
          <p></p>
          <button type="submit" className="btn btn-primary ">
            Wyślij
          </button>
        </form>
      </div>
    </>
  );
};
