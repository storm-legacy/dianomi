import React from 'react';
import { Link } from 'react-router-dom';
const UserDashboardPage = () => {
  const divItem = [
    {
      id: '1',
      name: 'wideo1',
      description: 'Lorens ipsum',
      author_id: 'Lorens',
    },
    {
      id: '2',
      name: 'wideo2',
      description: 'Lorens ipsum',
      author_id: 'Lorens',
    },
    {
      id: '3',
      name: 'wideo3',
      description: 'Lorens ipsum',
      author_id: 'Lorens',
    },
    {
      id: '4',
      name: 'wideo4',
      description: 'Lorens ipsum',
      author_id: 'Lorens',
    },
  ];
  return (
    <>
      <div className="dashbord  d-flex align-items-center">
        {divItem.map((item, index) => (
          <div key={index} className="card cardMY d-flex justify-content-center">
            <div className="p-2">
              <img src="..." className="card-img-top" alt="..." />
              <div className="card-body">
                <p className="card-text">
                  <h5>{item.name}</h5>
                  <p>{item.description}</p>
                  <p>{item.author_id}</p>
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default UserDashboardPage;
