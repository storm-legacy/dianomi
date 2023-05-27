import { createContext } from 'react';
import { User } from '../types/user.type';

interface AuthContext {
  user: User | null;
  setUser: (user: User | null) => void;
}

export const AuthContext = createContext<AuthContext>({
  user: null,
  setUser: () => {
    return null;
  },
});
