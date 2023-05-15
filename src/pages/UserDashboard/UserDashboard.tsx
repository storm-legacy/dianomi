import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import VideoService from '../../services/video.service';
const UserDashboardPage = () => {
  const divItem = [
    {
      id: '1',
      name: 'wideo1',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, id ',
      author_id: 'domino jachaś',
    },
    {
      id: '2',
      name: 'wideo2',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, id  ',
      author_id: 'Lorens',
    },
    {
      id: '3',
      name: 'wideo3',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris, ',
      author_id: 'Lorens',
    },
    {
      id: '4',
      name: 'wideo4',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi gravida massa mauris,  ',
      author_id: 'Lorens',
    },
  ];

  return (
    <div className="container">
      <div className="dashbord  d-flex align-items-center">
        {divItem.map((item, index) => (
          <Link to="/" key={index} className="card cardMY justify-content-center">
            <div className="p-2">
              <img src="./OIP.jpg" className="card-img-top" alt="logo kursu" />
              <div className="card-body">
                <div className="card-text">
                  <p className="lead">{item.name}</p>
                  <p>{item.description}</p>
                  <p>By:{item.author_id}</p>
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
