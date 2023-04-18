import { useRecoilState, useRecoilValue, useSetRecoilState } from 'recoil';
import { authAtom } from '../states/auth';
import axios from 'axios';

const backendURL = 'https://localhost/api/v1';

interface LoginProps {
  loginEmail: string;
  loginPassword: string;
  callback: () => void;
}

interface LogoutProps {
  callback: () => void;
}

const useAuthHelper = () => {
  const [auth, setAuth] = useRecoilState(authAtom);

  const login = ({ loginEmail, loginPassword, callback }: LoginProps) => {
    axios
      .post(`${backendURL}/auth/login`, {
        email: loginEmail,
        password: loginPassword,
      })
      .then((res) => {
        if (res.status === 200) {
          const accessToken = res.data['token'];
          localStorage.setItem('accessToken', accessToken);
          setAuth(accessToken);
          callback();
        }
      })
      .catch((err) => {
        console.error(err.response);
      });
  };

  // const refresh = () => {};

  const logout = ({ callback }: LogoutProps) => {
    console.log(`Bearer ${auth}`);
    axios
      .post(
        `${backendURL}/auth/logout`,
        {},
        {
          headers: {
            Authorization: `Bearer ${auth}`,
          },
        },
      )
      .then((res) => {
        if (res.status === 200) {
          localStorage.removeItem('accessToken');
          setAuth(null);
          callback();
        }
      })
      .catch((err) => {
        console.error(err.response);
      });
  };

  // const register = (username: string, password: string, passwordRepeat: string) => {};

  return {
    login,
    // register,
    // refresh,
    logout,
  };
};

export { useAuthHelper };
