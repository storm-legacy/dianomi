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
        console.log(err);
      });
  }, []);

  return (
    <div className="container">
      <div className="dashbord  d-flex align-items-center">
        {divItem.map((item, index) => (
          <Link to="/" key={index} className="card cardMY justify-content-center">
            <div className="p-2 myP">
              <img src={item.thumbnail_url} className="card-img-top myImg" alt="logo kursu" />
              <div className="card-body">
                <div className="card-text ">
                  <p className="lead">{item.name}</p>
                  <p className="myDes">{item.description}</p>
                </div>
              </div>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default UserDashboardPage;
