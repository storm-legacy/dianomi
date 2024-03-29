import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import videoService, { VideoPatchData } from '../../../services/video.service';
import { useNavigate } from 'react-router-dom';

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

export const VideoEdit = () => {
  const { VideoId } = useParams();
  const VideoIdInt = VideoId ? parseInt(VideoId, 10) : undefined;
  const [videoName, setVideoName] = useState('');
  const [isDisable, setIsDisabled] = useState(false);
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null);
  const [videoDiscription, setVideoDiscription] = useState('');
  const [videoTag, setVideoTag] = useState<string>('');
  const [defaultTagValue, setDefaultTagValue] = useState<string>('');
  const [isPremium, setIsPremium] = useState<boolean>(false);
  const navigate = useNavigate();
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
    request.then((res) => {
      console.log(res.data);
      setVideoTag(res.data.tags.toString());
      setVideoName(res.data.name);
      setVideoDiscription(res.data.description);
      setSelectedCategory(res.data.category_id);
      setDefaultTagValue(res.data.tags.toString());
      setIsPremium(res.data.IsPremium);
    });
  }, []);

  const handleSubmit = (event: any) => {
    event.preventDefault();
    let tagsArray: string[] = [defaultTagValue];

    console.log(tagsArray);
    if (defaultTagValue != videoTag) {
      tagsArray = videoTag
        .split(',')
        .map((value: string) => value.trim())
        .filter((value: string) => value.length >= 3 && value.length <= 12)
        .map((value: string) => value.toLowerCase());
    } else {
      setVideoTag(videoTag + ',');
      tagsArray = videoTag
        .split(',')
        .map((value: string) => value.trim())
        .filter((value: string) => value.length >= 3 && value.length <= 12)
        .map((value: string) => value.toLowerCase());
    }

    const videoData: VideoPatchData = {
      name: videoName,
      description: videoDiscription,
      category_id: selectedCategory,
      is_premium: isPremium,
      Tags: tagsArray,
    };
    const { request } = videoService.editVideo(VideoIdInt, videoData);
    request
      .then((ress) => {
        console.log(ress);
        setIsDisabled(true);
        navigate('/VideoList');
      })
      .catch((err) => {
        console.error(err.message);
        console.log(isPremium);
      });
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
              maxLength={100}
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
              maxLength={500}
              style={{ height: '15dvh' }}
              id="exampleFormControlTextarea1"
              value={videoDiscription}
              onChange={(event) => setVideoDiscription(event.target.value)}
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
            <input
              className="form-control"
              type="text"
              value={videoTag}
              onChange={(event) => setVideoTag(event.target.value)}
            />
          </label>
          <p></p>
          <label>
            <input
              className="form-check-input"
              type="checkbox"
              checked={isPremium}
              onChange={(e) => setIsPremium(e.target.checked)}
            />{' '}
            Premium material
          </label>
          <p></p>
          <button type="submit" className="btn btn-primary " disabled={isDisable}>
            Send
          </button>
        </form>
      </div>
    </>
  );
};
