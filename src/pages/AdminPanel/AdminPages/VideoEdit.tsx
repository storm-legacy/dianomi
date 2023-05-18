import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import videoService from '../../../services/video.service';
export const VideoEdit = () => {
  interface VideoItemData {
    id: number;
    name: string;
    description: string;
    category_id: number;
    tags: string[];
  }
  const [videoListItem, setVideoListItem] = useState<VideoItemData[]>([]);
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;

  const [videoName, setVideoName] = useState('');
  const [videoFile, setVideoFile] = useState('');
  const [videoCategori, setVideoCategori] = useState('');
  const [videoDiscription, setVideoDiscription] = useState('');
  const [videoTag, setVideoTag] = useState('');
  useEffect(() => {
    const { request } = videoService.takeVideoId(VideoIdInt);
    request
      .then((res) => {
        console.log(res);
        const Videodata = res.data.map(
          (Videodata: { id: number; name: string; description: string; category_id: number; tags: string[] }) => {
            return {
              id: Videodata.id,
              name: Videodata.name,
              description: Videodata.description,
              category: Videodata.category_id,
              tags: Videodata.tags,
            };
          },
        );
        setVideoListItem(Videodata);
        console.log(Videodata);
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  const handleSubmit = (event: any) => {
    event.preventDefault();

    const videoData = {
      videoName: videoName,
      videoDiscription: videoDiscription,
      videoTag: videoTag,
      videoCategori: videoCategori,
      videoFile: videoFile,
    };
  };
  return (
    <div>
      {videoListItem.map((item) => {
        if (item.id === VideoIdInt) {
          return (
            <div
              key={item.id}
              className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded"
            >
              <h3>Edit Video</h3>
              <p></p>
              <form onSubmit={handleSubmit} className="row">
                <label>
                  <p>Name</p>
                  <input
                    className="form-control"
                    type="text"
                    value={item.name}
                    onChange={(event) => setVideoName(event.target.value)}
                  />
                </label>
                <label>
                  <p>Discription</p>
                  <input
                    className="form-control"
                    type="text"
                    value={item.description}
                    onChange={(event) => setVideoDiscription(event.target.value)}
                  />
                </label>
                <label>
                  <p>File</p>
                  <input
                    className="form-control"
                    type="file"
                    value={videoFile}
                    onChange={(event) => setVideoFile(event.target.value)}
                  />
                </label>
                <label>
                  <p>Category</p>
                  <input
                    className="form-control"
                    type="list"
                    list="CategoryList"
                    value={item.category_id}
                    onChange={(event) => setVideoCategori(event.target.value)}
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
                    value={item.tags}
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
                  Edit
                </button>
              </form>
            </div>
          );
        } else {
          return null;
        }
      })}
    </div>
  );
};
