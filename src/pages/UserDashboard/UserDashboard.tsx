import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import videoService from '../../services/video.service';

const UserDashboardPage = () => {
  interface VideoItemData {
    id: number;
    name: string;
    description: string;
    category: string;
    tags: string[];
    thumbnail_url: string;
  }
  const [divItem, setDivItem] = useState<VideoItemData[]>([]);
  useEffect(() => {
    const { request } = videoService.takeVideoRecommended();
    request
      .then((res) => {
        console.log(res);
        const Videodata = res.data.map(
          (Videodata: {
            id: number;
            name: string;
            description: string;
            category: string;
            tags: string[];
            thumbnail_url: string;
          }) => {
            return {
              id: Videodata.id,
              name: Videodata.name,
              description: Videodata.description,
              category: Videodata.category,
              tags: Videodata.tags,
              thumbnail_url: Videodata.thumbnail_url,
            };
          },
        );
        setDivItem(Videodata);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  return (
    <>
      <div className="container m-0 p-4">
        <div className="row row-cols-3">
          {divItem ? (
            divItem.map((item, index) => (
              <div className="col" key={index}>
                <Link to={'/VideoPlayer/' + item.id} className="card cardMY justify-content-center">
                  <div className="p-2 myP">
                    <img
                      src={'http://localhost:9000/thumbnails/' + item.thumbnail_url}
                      className="card-img-top myImg"
                      alt="logo kursu"
                    />
                    <div className="card-body">
                      <div className="card-text">
                        <p className="lead">{item.name}</p>
                        <p className="myDes">{`${item.description.substring(0, 124)}...`}</p>
                      </div>
                    </div>
                  </div>
                </Link>
              </div>
            ))
          ) : (
            <span>No videos to show</span>
          )}
        </div>
      </div>
    </>
  );
};

export default UserDashboardPage;
