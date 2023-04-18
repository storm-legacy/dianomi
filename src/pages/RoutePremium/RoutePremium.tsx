import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const baseURL = 'https://localhost/api/v1';

const RoutePremium = () => {
  const navigate = useNavigate();
  const accessToken = localStorage.getItem('accessToken');
  const header = {
    'Content-Type': 'application/json',
    Authorization: `Bearer ${accessToken}`,
  };
  const [data, setData] = useState('');

  useEffect(() => {
    axios
      .get(`${baseURL}/routePremium`, {
        headers: header,
      })
      .then((response) => {
        setData(response.data.status);
      })
      .catch((error) => {
        if ([401, 403].includes(error.response.status)) {
          localStorage.removeItem('accessToken');
          navigate('/');
        }
        console.log(error);
      });
  }, [data]);

  return <>{data}</>;
};

export default RoutePremium;
