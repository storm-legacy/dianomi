import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { Upload } from '@aws-sdk/lib-storage';
import { S3Client, S3 } from '@aws-sdk/client-s3';
import VideoService, { VideoAddData } from '../../../services/video.service';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../../../context/AuthContext';

const s3endpoint = 'http://localhost:9000';

export const VideoAdd = () => {
  const [videoArrLong, SetVideoArrLong] = useState<number>(0);
  const [videoName, setVideoName] = useState('');
  const [file, setFile] = useState<File>();
  const [thumbnailFile, setThumbnailFile] = useState<File>();
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null);
  const [videoDescription, setVideoDescription] = useState('');
  const [tags, setTags] = useState<string>('');
  const [width, setWidth] = useState(0);
  const [width2, setWidth2] = useState(0);
  const [isDisabled, setIsDisabled] = useState<boolean>(false);
  const navigate = useNavigate();

  interface Category {
    ID: number;
    Name: string;
  }
  const [categoriesArr, setCategoriesArr] = useState<Category[]>([]);
  const { user } = useContext(AuthContext);

  const ifAdd = () => {
    const { request } = VideoService.takeVideo();
    request.then((res) => {
      console.log('sprawdzam');
      if (videoArrLong != res.data.length) {
        console.log('zmieniło sie');
        navigate('/');
      } else {
        console.log(videoArrLong);
        setTimeout(ifAdd, 25000);
      }
    });
  };

  useEffect(() => {
    const { request } = VideoService.takeVideo();
    request.then((res) => {
      console.log(res.data.length);
      SetVideoArrLong(res.data.length);
    });
  }, []);

  useEffect(() => {
    if (file == null || thumbnailFile == null) {
      setIsDisabled(true);
    } else {
      setIsDisabled(false);
    }
  }, [file, thumbnailFile]);

  useEffect(() => {
    const { request } = VideoService.takeCategori();
    request
      .then((res) => {
        console.log(res);
        const categories = res.data.map((category: { ID: any; Name: any }) => {
          return {
            ID: category.ID,
            Name: category.Name,
          };
        });
        setCategoriesArr(categories);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    setIsDisabled(true);

    const data = {
      Action: 'AssumeRoleWithCustomToken',
      Token: user?.authToken,
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
          Token: user?.authToken,
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

      const thumbnailUploads3 = new Upload({
        client: new S3(s3config) || new S3Client(s3config),
        queueSize: 4,
        leavePartsOnError: false,
        params: {
          ContentType: thumbnailFile?.type,
          Bucket: 'uploads',
          Key: thumbnailFile?.name,
          Body: thumbnailFile,
        },
      });

      thumbnailUploads3.on('httpUploadProgress', (progress) => {
        if (progress.total && progress.loaded) {
          const percentage = Math.round((progress.loaded / progress.total) * 100);
          setWidth2(percentage);
        }
      });
      await thumbnailUploads3.done();

      const tagsArray = tags
        .split(',')
        .map((value: string) => value.trim())
        .filter((value: string) => value.length >= 3 && value.length <= 12)
        .map((value: string) => value.toLowerCase());

      // Send data to backend
      const data: VideoAddData = {
        name: videoName,
        description: videoDescription,
        file_name: String(file?.name),
        thumbnail_name: String(thumbnailFile?.name),
        file_bucket: 'uploads',
        category_id: selectedCategory,
        tags: tagsArray,
      };

      const { request } = VideoService.sendVideo(data);
      request
        .then((res) => {
          console.log(res);
          ifAdd();
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
            <textarea
              className="form-control"
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
            <p>Thumbnail</p>
            <input
              className="form-control"
              id="thumbnailfilefield"
              type="file"
              onChange={(e) => setThumbnailFile(e.target?.files?.[0])}
            />
            {
              <div className="progress">
                <div
                  className="progress-bar progress-bar-striped progress-bar-animated"
                  role="progressbar"
                  style={{ width: `${width2}%` }}
                ></div>
              </div>
            }
          </label>
          <label>
            <p>Categories</p>
            <select
              className="form-select"
              defaultValue={'DEFAULT'}
              onChange={(event) => setSelectedCategory(parseInt(event.target.value))}
            >
              <option value="DEFAULT">Please select option</option>
              {categoriesArr.map((item) => (
                <option value={item.ID} key={item.ID}>
                  {' '}
                  {item.Name}
                </option>
              ))}
            </select>
          </label>
          <label>
            <p>Tag</p>
            <input className="form-control" type="text" value={tags} onChange={(e) => setTags(e.target.value)} />
          </label>
          <p></p>
          <button type="submit" className="btn btn-primary " disabled={isDisabled}>
            Send
          </button>
        </form>
      </div>
    </>
  );
};
