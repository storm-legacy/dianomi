import React, { useContext, useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import profileService, { Email } from '../../services/profile.service';
import { AuthContext } from '../../context/AuthContext';
import { TbCrown } from 'react-icons/tb';

export const UserHistory = () => {
  const { user } = useContext(AuthContext);
  const data: Email = {
    email: user?.email,
  };
  interface MetricData {
    video_id: number;
    name: string;
    description: string;
    IsPremium: boolean;
    thumbnail_url: string;
    updated_at: string[];
  }
  const [metric, setMetric] = useState<MetricData[]>([]);
  useEffect(() => {
    const { request } = profileService.GetUserVideoMetric(data);
    request
      .then((res) => {
        const Data = res.data.map(
          (Data: {
            video_id: number;
            name: string;
            description: string;
            IsPremium: boolean;
            thumbnail_url: string;
            updated_at: string[];
          }) => {
            return {
              video_id: Data.video_id,
              name: Data.name,
              description: Data.description,
              IsPremium: Data.IsPremium,
              thumbnail_url: Data.thumbnail_url,
              updated_at: Data.updated_at,
            };
          },
        );
        setMetric(Data);
        console.log(Data.updated_at);
      })
      .catch((err) => console.error(err));
  }, []);
  return (
    <div className="container mt-5 p-5 shadow-lg">
      {metric
        ? metric.map((item, index) => (
            <Link
              className={`row ${item.IsPremium ? 'border border-warning' : ''}`}
              key={index}
              to={
                item.IsPremium
                  ? user?.role != 'free'
                    ? '/VideoPlayer/' + item.video_id
                    : ''
                  : '/VideoPlayer/' + item.video_id
              }
            >
              <div className="col-3 border">
                {' '}
                <img src={item.thumbnail_url} className="card-img-top myImg" alt="logo kursu" />{' '}
              </div>
              <div className="col-6 border">
                <p className="h4">
                  {' '}
                  {String(item.name)}
                  {item.IsPremium && <TbCrown style={{ color: '#DAA520' }} />}
                </p>
                <p>{`${item.description.substring(0, 324)}...`}</p>
              </div>
              <div className="col-3 border">
                {' '}
                <p>Last Watched</p>
                <p>{String(item.updated_at[0])}</p>
              </div>
            </Link>
          ))
        : ''}
    </div>
  );
};
