import React from 'react';
import { FiUser } from 'react-icons/fi';

export const UserHistory = () => {
  return (
    <div className="container mt-5 p-5 shadow-lg">
      <div className="row">
        <div className="col-3 border">
          {' '}
          <FiUser size={200}></FiUser>
        </div>
        <div className="col-9 border">
          <p className="h4">Lorem ipsum</p>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis et lectus et mauris viverra varius in eget
            nulla. Morbi varius magna augue, et congue orci posuere ut.{' '}
          </p>
          <p>JAVA</p>
          <p>Tag1,tag2</p>
        </div>
      </div>
    </div>
  );
};
