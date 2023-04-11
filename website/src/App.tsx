import { Fragment } from 'react';
import { CssBaseline } from '@mui/material';
import { BrowserRouter } from 'react-router-dom';
import { atom, useRecoilState } from 'recoil';

const loggedState = atom<boolean>({
  key: 'loggedState',
  default: true,
});

type Actions = {
  toggle: () => void;
  signOut: () => void;
  signIn: () => void;
};

function useLoginStatus(): [boolean, Actions] {
  const [isLogged, setIsLogged] = useRecoilState(loggedState);

  function toggle() {
    setIsLogged((isLogged: boolean) => !isLogged);
  }

  function signOut() {
    setIsLogged(false);
  }

  function signIn() {
    setIsLogged(true);
  }

  return [isLogged, { toggle, signOut, signIn }];
}

function App() {
  const [isLogged, loginActions] = useLoginStatus();
  return (
    <Fragment>
      <CssBaseline />
      {isLogged ? (
        <img
          src="https://i.redd.it/ayih4qogh2a51.png"
          onClick={loginActions.signOut}
          alt="nice bearo"
        />
      ) : (
        <div />
      )}
      <BrowserRouter></BrowserRouter>
    </Fragment>
  );
}

export default App;
