import axios, { CanceledError } from 'axios';

const createAxiosInstance = () => {
  const http = axios.create({ baseURL: 'https://localhost/api/v1' });

  const getAuthorizationHeader = () => {
    const token = String(localStorage.getItem('token'));
    return token ? `Bearer ${token}` : '';
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
        localStorage.removeItem('token');
        delete http.defaults.headers.common['Authorization'];
      }
      return Promise.reject(error);
    },
  );

  return http;
};

export default createAxiosInstance();
