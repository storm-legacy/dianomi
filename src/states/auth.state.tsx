import { atom } from 'recoil';

interface CurrentUser {
  name: string;
  email: string;
  token: string;
}

const currentUserAtom = atom({
  key: 'current_user',
  default: localStorage.getItem('currentUser')
    ? (JSON.parse(String(localStorage.getItem('currentUser'))) as CurrentUser)
    : ({} as CurrentUser),
});

export { currentUserAtom };
export type { CurrentUser };
