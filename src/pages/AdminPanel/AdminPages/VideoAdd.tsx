import React, { useState } from 'react';
import env from '../../../environment/env';
import axios from 'axios';

export const VideoAdd = () => {
  const [videoName, setVideoName] = useState('');
  const [videoFile, setVideoFile] = useState('');
  const [videoAuthor, setVideoAuthor] = useState('');
  const [videoCategori, setVideoCategori] = useState('');
  const [videoDiscription, setVideoDiscription] = useState('');
  const [videoTag, setVideoTag] = useState('');

  const handleSubmit = (event: any) => {
    event.preventDefault();

    const videoData = {
      videoName: videoName,
      videoDiscription: videoDiscription,
      videoAuthor: videoAuthor,
      videoTag: videoTag,
      videoCategori: videoCategori,
      videoFile: videoFile,
    };
    console.log({ videoData });
  };

  // const sendVideoHandler = async () => {
  //   // Get token for S3
  //   const SecretAccessKey = '';
  //   const SessionToken = '';

  //   const data = {
  //     Action: 'AssumeRoleWithCustomToken',
  //     Token: localStorage.getItem('accessToken'),
  //     Version: '2011-06-05',
  //     DurationSeconds: '3600',
  //     RoleArn: 'arn:minio:iam:::role/idmp-dianomi-server-auth',
  //   };

  //   console.log(data);

  //   // Get permissions from s3
  //   axios
  //     .post(`${env.S3_URL}`, data)
  //     .then((res) => {
  //       if (res.status === 200) {
  //         console.log(res);
  //       }
  //     })
  //     .catch((err) => {
  //       console.log(err);
  //     });
  // };

  return (
    <>
      <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
        <h3>Dodaj Wideo</h3>
        <p></p>
        <form onSubmit={handleSubmit} className="row">
          <label>
            <p>Nazwa</p>
            <input
              className="form-control"
              type="text"
              value={videoName}
              onChange={(event) => setVideoName(event.target.value)}
            />
          </label>
          <label>
            <p>Opis</p>
            <input
              className="form-control"
              type="text"
              value={videoDiscription}
              onChange={(event) => setVideoDiscription(event.target.value)}
            />
          </label>
          <label>
            <p>Plik</p>
            <input
              className="form-control"
              type="file"
              value={videoFile}
              onChange={(event) => setVideoFile(event.target.value)}
            />
          </label>
          <label>
            <p>Autor</p>
            <input
              className="form-control"
              type="text"
              value={videoAuthor}
              onChange={(event) => setVideoAuthor(event.target.value)}
            />
          </label>
          <label>
            <p>Kategoria</p>
            <input
              className="form-control"
              type="list"
              list="CategoryList"
              value={videoCategori}
              onChange={(event) => setVideoCategori(event.target.value)}
            />
            <datalist id="CategoryList">
              <option value="C++" />
              <option value="Go" />
              <option value="HowTo" />
            </datalist>
          </label>
          <label>
            <p>tag</p>
            <input
              className="form-control"
              list="TagList"
              value={videoTag}
              onChange={(event) => setVideoTag(event.target.value)}
            />
            <datalist id="TagList">
              <option value="1" />
              <option value="2" />
              <option value="3" />
            </datalist>
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
