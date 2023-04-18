import { atom } from 'recoil';

const authAtom = atom({
  key: 'auth',
  default: localStorage.getItem('accessToken'),
});

export { authAtom };
