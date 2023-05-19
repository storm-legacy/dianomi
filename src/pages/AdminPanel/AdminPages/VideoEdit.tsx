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
  interface Category {
    ID: number;
    Name: string;
  }
  const [CatrgorisId, setCatrgorisId] = useState('');
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  const [videoName, setVideoName] = useState('');
  const [videoFile, setVideoFile] = useState('');
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null);
  const [videoDiscription, setVideoDiscription] = useState('');
  const [videoTag, setVideoTag] = useState('');

  const [categoriesArr, setCategoriesArr] = useState<Category[]>([]);
  useEffect(() => {
    const { request } = videoService.takeCategori();
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

  useEffect(() => {
    const { request } = videoService.takeVideoId(VideoIdInt);
    request
      .then((res) => {
        setVideoTag(res.data.tags);
        setVideoName(res.data.name);
        setVideoDiscription(res.data.description);
        setSelectedCategory(res.data.category_id);
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
      videoCategori: selectedCategory,
      videoFile: videoFile,
    };
  };
  return (
    <>
      <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
        <h3>Video Edit</h3>
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
              style={{ height: '15dvh' }}
              id="exampleFormControlTextarea1"
              value={videoDiscription}
            ></textarea>
          </label>

          <label>
            <p>Categories</p>
            <select className="form-select" onChange={(event) => setSelectedCategory(parseInt(event.target.value))}>
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
            <input className="form-control" type="text" value={videoTag} />
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
