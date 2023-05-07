import http from './http';
import { CurrentUser } from '../states/auth.state';
import { useRecoilValue } from 'recoil';
import { currentUserAtom } from '../states/auth.state';
import env from '../environment/env';

interface LoginData {
  status: string;
  data: CurrentUser;
}

interface RefreshData {
  status: string;
  data: {
    token: string;
  };
}

interface AuthResponse {
  status: string;
  data: string;
}

class AuthService {
  login(loginEmail: string, loginPassword: string) {
    return http.post<LoginData>('/auth/login', { email: loginEmail, password: loginPassword });
  }

  register(regEmail: string, regPassword: string, regPasswordRepeat: string) {
    return http.post<AuthResponse>('/auth/register', {
      email: regEmail,
      password: regPassword,
      password_repeat: regPasswordRepeat,
    });
  }

  refresh() {
    return http.post<RefreshData>('/auth/refresh');
  }

  logout() {
    return http.post<AuthResponse>('/auth/logout');
  }
}

export default new AuthService();
export type { LoginData, AuthResponse, RefreshData };
