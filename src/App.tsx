import React from 'react';
import { BrowserRouter, Routes, Route, Link, useNavigate } from 'react-router-dom';
import { useRecoilValue } from 'recoil';

import Protected from './components/Protected';
import LoginPage from './pages/LoginPage/LoginPage';
import RegisterPage from './pages/RegisterPage/RegisterPage';
import NotFoundPage from './pages/NotFound/NotFound';

import RouteNormal from './pages/RouteNormal/RouteNormal';
import RoutePremium from './pages/RoutePremium/RoutePremium';
import RouteAdmin from './pages/RouteAdmin/RouteAdmin';

import { authAtom } from './states/auth';

import './App.css';
import { useAuthHelper } from './helpers/authHelper';
import UserDashboardPage from './pages/UserDashboard/UserDashboard';

function App() {
  const authHelper = useAuthHelper();
  const auth = useRecoilValue(authAtom);
  const isLogged = auth ? true : false;

  return (
    <>
      <div className="container">
        <BrowserRouter>
          <Routes>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route
              path="/"
              element={
                <Protected isLoggedIn={isLogged}>
                  <ul>
                    <li>
                      <Link to="/routeNormal">Normal user route</Link>
                    </li>
                    <li>
                      <Link to="/routePremium">Premium user route</Link>
                    </li>
                    <li>
                      <Link to="/routeAdmin">Admin user route</Link>
                    </li>
                    <br />
                    <button
                      onClick={() => {
                        authHelper.logout({
                          callback: () => {
                            window.location.reload();
                          },
                        });
                      }}
                    >
                      Logout
                    </button>
                  </ul>
                </Protected>
              }
            />
            <Route
              path="/routeNormal"
              element={
                <Protected isLoggedIn={isLogged}>
                  <RouteNormal />
                </Protected>
              }
            />
            <Route
              path="/routePremium"
              element={
                <Protected isLoggedIn={isLogged}>
                  <RoutePremium />
                </Protected>
              }
            />
            <Route
              path="/routeAdmin"
              element={
                <Protected isLoggedIn={isLogged}>
                  <RouteAdmin />
                </Protected>
              }
            />
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
