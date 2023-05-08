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

import { ProfilePage } from './pages/ProfilePage/ProfilePage';
import SidePanel from './pages/SidePanel/SidePanel';
import { AdminPanel } from './pages/AdminPanel/AdminPanel';
import { VideoAdd } from './pages/AdminPanel/AdminPages/VideoAdd';
import UserList from './pages/AdminPanel/AdminPages/UserList';
import UserEdit from './pages/AdminPanel/AdminPages/UserEdit';

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
                  <SidePanel />
                  <UserDashboardPage />
                </Protected>
              }
            />
            <Route
              path="/UserDashbord"
              element={
                <Protected isLoggedIn={isLogged}>
                  <SidePanel />
                  <UserDashboardPage />
                </Protected>
              }
            />

            <Route
              path="/ProfilePage"
              element={
                <Protected isLoggedIn={isLogged}>
                  <SidePanel />
                  <ProfilePage />
                </Protected>
              }
            />
            <Route
              path="/UserList"
              element={
                <Protected isLoggedIn={isLogged}>
                  <SidePanel />
                  <UserList />
                </Protected>
              }
            />
            <Route
              path="/UserEdit/:UserId"
              element={
                <Protected isLoggedIn={isLogged}>
                  <SidePanel />
                  <UserEdit />
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
            <Route
              path="/AdminPanel"
              element={
                <Protected isLoggedIn={isLogged}>
                  <SidePanel />
                  <AdminPanel />
                </Protected>
              }
            />
            <Route
              path="/VideoAdd"
              element={
                <Protected isLoggedIn={isLogged}>
                  <SidePanel />
                  <VideoAdd />
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
