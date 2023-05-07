import axios from 'axios';
import env from '../environment/env';
import { CurrentUser } from '../states/auth.state';

export default axios.create({
  baseURL: env.BACKEND_URL,
  headers: {
    'Content-Type': 'application/json',
    Authorization: localStorage.getItem('currentUser')
      ? `Bearer ${(JSON.parse(String(localStorage.getItem('currentUser'))) as CurrentUser).token}`
      : '',
  },
});
