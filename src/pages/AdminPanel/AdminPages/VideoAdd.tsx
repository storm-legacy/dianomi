import React, { useState } from 'react';
import axios from 'axios';
import { Upload } from '@aws-sdk/lib-storage';
import { S3Client, S3 } from '@aws-sdk/client-s3';
import adminService, { VideoAddData } from '../../../services/admin.service';

const s3endpoint = 'http://localhost:9000';

export const VideoAdd = () => {
  const [videoName, setVideoName] = useState('');
  const [file, setFile] = useState<File>();
  const [videoCategory, setVideoCategory] = useState('');
  const [videoDescription, setVideoDescription] = useState('');
  const [tag, setTag] = useState('');
  const [videoTag, setVideoTag] = useState(['']);
  const [width, setWidth] = useState(0);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    const inputValue = tag;
    const newValues = inputValue
      .split(',')
      .map((value) => value.trim())
      .filter((value) => value.length >= 3 && value.length <= 12)
      .map((value) => value.toLowerCase());
    setVideoTag(newValues);

    const data = {
      Action: 'AssumeRoleWithCustomToken',
      Token: localStorage.getItem('token'),
      Version: '2011-06-15',
      DurationSeconds: 3600,
      RoleArn: 'arn:minio:iam:::role/idmp-dianomi-server-auth',
    };

    try {
      const res = await axios({
        method: 'post',
        url: s3endpoint,
        params: {
          Action: 'AssumeRoleWithCustomToken',
          Token: String(localStorage.getItem('token')),
          Version: '2011-06-15',
          DurationSeconds: '3600',
          RoleArn: 'arn:minio:iam:::role/idmp-dianomi-server-auth',
        },
      });

      const parser = new DOMParser();
      const xml = parser.parseFromString(res.data, 'application/xml');
      const accessKeyId = String(xml.querySelector('AccessKeyId')?.textContent);
      const secretAccessKey = String(xml.querySelector('SecretAccessKey')?.textContent);
      const sessionToken = String(xml.querySelector('SessionToken')?.textContent);

      if (!accessKeyId || !secretAccessKey || !sessionToken) {
        console.error('Access credentials empty');
        return;
      }

      const s3config = {
        region: 'us-east-1',
        endpoint: 'http://localhost:9000',
        credentials: {
          accessKeyId: accessKeyId,
          secretAccessKey: secretAccessKey,
          sessionToken: sessionToken,
        },
        forcePathStyle: true,
      };

      const parallelUploads3 = new Upload({
        client: new S3(s3config) || new S3Client(s3config),
        queueSize: 4,
        leavePartsOnError: false,
        params: {
          ContentType: file?.type,
          Bucket: 'uploads',
          Key: file?.name,
          Body: file,
        },
      });

      parallelUploads3.on('httpUploadProgress', (progress) => {
        if (progress.total && progress.loaded) {
          const percentage = Math.round((progress.loaded / progress.total) * 100);
          setWidth(percentage);
        }
      });

      await parallelUploads3.done();

      // Send data to backend
      const data: VideoAddData = {
        name: videoName,
        description: videoDescription,
        file_name: String(file?.name),
        file_bucket: 'uploads',
        category_id: null,
        tags: videoTag,
      };

      const { request } = adminService.sendVideo(data);
      request
        .then((res) => {
          console.log(res);
        })
        .catch((err) => {
          console.log(err);
        });
    } catch (err) {
      console.error(err);
    }
  };
  return (
    <>
      <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
        <h3>Add Video</h3>
        <p></p>
        <form onSubmit={handleSubmit} className="row">
          <label>
            <p>Title</p>
            <input
              className="form-control"
              type="text"
              value={videoName}
              onChange={(event) => setVideoName(event.target.value)}
            />
          </label>
          <label>
            <p>Description</p>
            <input
              className="form-control"
              type="text"
              value={videoDescription}
              onChange={(event) => setVideoDescription(event.target.value)}
            />
          </label>
          <label>
            <p>File</p>
            <input
              className="form-control"
              id="filefield"
              type="file"
              onChange={(e) => setFile(e.target?.files?.[0])}
            />
            {
              <div className="progress">
                <div
                  className="progress-bar progress-bar-striped progress-bar-animated"
                  role="progressbar"
                  style={{ width: `${width}%` }}
                ></div>
              </div>
            }
          </label>
          <label>
            <p>Categories</p>
            <input
              className="form-control"
              type="list"
              list="CategoryList"
              value={videoCategory}
              onChange={(event) => setVideoCategory(event.target.value)}
            />
            <datalist id="CategoryList">
              <option value="C++" />
              <option value="Go" />
              <option value="HowTo" />
            </datalist>
          </label>
          <label>
            <p>Tag</p>
            <input
              className="form-control"
              list="TagList"
              value={tag}
              onChange={(event) => setTag(event.target.value)}
            />
            <datalist id="TagList">
              {' '}
              <option value="1" />
              <option value="2" />
              <option value="3" />
            </datalist>
          </label>
          <p></p>
          <button type="submit" className="btn btn-primary ">
            Send
          </button>
        </form>
      </div>
    </>
  );
};
