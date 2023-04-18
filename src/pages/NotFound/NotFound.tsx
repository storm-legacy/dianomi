import React from 'react';
import { Link } from 'react-router-dom';

const NotFoundPage = () => {
  return (
    <>
      <div className="container d-flex justify-content-center align-items-center">
        <h1>Page not found</h1>
        <Link to="/">Return</Link>
      </div>
    </>
  );
};

export default NotFoundPage;
