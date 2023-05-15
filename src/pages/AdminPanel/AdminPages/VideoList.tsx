import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import videoService from '../../../services/video.service';
export const VideoList = () => {
  interface VideoItemData {
    id: number;
    name: string;
    description: string;
    category: string;
    tags: string[];
  }
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
        console.log(err);
      });
  }, []);

  return (
    <>
      <div className="text-center">
        <div className=" position-absolute top-50 start-50 translate-middle">
          <div className="row myRow myRowCont overflow-auto">
            <h1>Video List</h1>
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
                  Delete <Link to={'/VideoEdit/' + item.id}>Edit</Link> Ban
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};
