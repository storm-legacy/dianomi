import axios, { CanceledError } from 'axios';
import { User } from '../types/user.type';

const createAxiosInstance = () => {
  const http = axios.create({ baseURL: 'https://localhost/api/v1' });

  const getAuthorizationHeader = () => {
    const user: User = JSON.parse(String(localStorage.getItem('user')));
    return user ? `Bearer ${user.authToken}` : '';
  };

  http.interceptors.request.use(
    (config) => {
      config.headers.Authorization = getAuthorizationHeader();
      return config;
    },
    (error) => {
      return error;
    },
  );

  http.interceptors.response.use(
    async (response) => {
      return response;
    },
    async (error) => {
      if (error instanceof CanceledError) {
        return Promise.reject(error);
      }
      if ([401, 403].includes(error.response.status)) {
        localStorage.clear();
        delete http.defaults.headers.common['Authorization'];
      }
      return Promise.reject(error);
    },
  );

  return http;
};

export default createAxiosInstance();
