import axios, { CanceledError } from 'axios';

const token = String(localStorage.getItem('token'));
const http = axios.create({ baseURL: 'https://localhost/api/v1' });

http.defaults.headers.common['Authorization'] = token ? `Bearer ${token}` : '';
http.defaults.headers.common['Accept'] = 'application/json';
http.defaults.headers.common['Content-Type'] = 'application/json';

http.interceptors.response.use(
  async (response) => {
    return response;
  },
  async (error) => {
    if (!(error instanceof CanceledError) && [401, 403].includes(error.response.status)) {
      localStorage.removeItem('token');
      http.defaults.headers.common['Authorization'] = '';
    }
    return Promise.reject(error);
  },
);

export default http;
