import { atom } from 'recoil';

const authAtom = atom({
  key: 'auth',
  default: JSON.parse(String(localStorage.getItem('user'))),
});

export { authAtom };
