import React, { useEffect, useState } from 'react';
import videoService from '../../../services/video.service';
import { Report } from 'notiflix';
import { Link } from 'react-router-dom';

interface VideoItemData {
  id: number;
  name: string;
  description: string;
  category: string;
  tags: string[];
}

export const VideoList = () => {
  const [videoListItem, setVideoListItem] = useState<VideoItemData[]>([]);

  useEffect(() => {
    const { request } = videoService.takeVideo();
    request
      .then((res) => {
        console.log(res);
        const Videodata = res.data.map(
          (Videodata: { id: number; name: string; description: string; category: string; tags: string[] }) => {
            return {
              id: Videodata.id,
              name: Videodata.name,
              description: Videodata.description,
              category: Videodata.category,
              tags: Videodata.tags,
            };
          },
        );
        setVideoListItem(Videodata);
      })
      .catch((err) => {
        Report.failure(
          'Problem with fetching videos',
          `Video could not be fetched from the server. Message: ${err.message}`,
          'Okay',
        );
      });
  }, []);

  const deleteVideo = (value: number) => {
    console.log(value);
    const { request } = videoService.deleteVideo(value);
    request.then((res) => window.location.reload()).catch((err) => console.error(err.message));
  };

  return (
    <>
      <div className="px-2 text-center">
        <h1>Video List</h1>
        <div className="row myRow myRowCont overflow-auto">
          <div className="row myRow">
            <div className="col-1 border border-primary">ID</div>
            <div className="col-2 border border-primary">Name</div>
            <div className="col-3 border border-primary">Description</div>
            <div className="col-2  border border-primary">Tag</div>
            <div className="col-2  border border-primary">Category</div>
            <div className="col-2 border border-primary"> Settings</div>
          </div>
          {videoListItem.map((item, index) => (
            <div className="row" key={index}>
              <div className="col-1 border border-primary">{item.id}</div>
              <div className="col-2 border border-primary">{item.name}</div>
              <div className="col-3 border border-primary">{item.description}</div>
              <div className="col-2  border border-primary">{item.tags}</div>
              <div className="col-2  border border-primary">{item.category}</div>
              <div className="col-2 border border-primary">
                {' '}
                <button className="custom-link link-primary" onClick={() => deleteVideo(item.id)}>
                  {' '}
                  Delete{' '}
                </button>{' '}
                <Link to={'/VideoEdit/' + item.id}>Edit</Link>
              </div>
            </div>
          ))}
        </div>
      </div>
    </>
  );
};
