import React, { useEffect } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import authService from './services/auth.service';

import Protected from './components/Protected';
import ProtectedAdmin from './components/ProtectedAdmin';
import ProtectedAuth from './components/ProtectedAuth';
import LoginPage from './pages/LoginPage/LoginPage';
import RegisterPage from './pages/RegisterPage/RegisterPage';
import NotFoundPage from './pages/NotFound/NotFound';

import UserDashboardPage from './pages/UserDashboard/UserDashboard';

import { ProfilePage } from './pages/ProfilePage/ProfilePage';
import SidePanel from './pages/SidePanel/SidePanel';
import { AdminPanel } from './pages/AdminPanel/AdminPanel';
import { VideoAdd } from './pages/AdminPanel/AdminPages/VideoAdd';
import UserList from './pages/AdminPanel/AdminPages/UserList';
import UserEdit from './pages/AdminPanel/AdminPages/UserEdit';
import { VideoList } from './pages/AdminPanel/AdminPages/VideoList';
import { VideoEdit } from './pages/AdminPanel/AdminPages/VideoEdit';

import './App.css';
import { CanceledError } from 'axios';
import CategoriAdd from './pages/AdminPanel/AdminPages/CategoriAdd';
import { PasswordResPage } from './pages/LoginPage/PasswordResPage';
import { CodePage } from './pages/LoginPage/CodePage';

function App() {
  useEffect(() => {
    const { request, cancel } = authService.connectionCheck();
    request
      .then(() => {
        console.info('Authorized');
      })
      .catch((err) => {
        if (err instanceof CanceledError) return;
        console.warn('Unauthorized');
      });

    return () => cancel();
  }, []);

  return (
    <>
      <div className="container">
        <BrowserRouter>
          <Routes>
            <Route
              path="/login"
              element={
                <ProtectedAuth>
                  <LoginPage />
                </ProtectedAuth>
              }
            />
            <Route
              path="/register"
              element={
                <ProtectedAuth>
                  <RegisterPage />
                </ProtectedAuth>
              }
            />
            <Route
              path="/verificate"
              element={
                <ProtectedAuth>
                  <CodePage />
                </ProtectedAuth>
              }
            />
            <Route
              path="/password/reset"
              element={
                <ProtectedAuth>
                  <PasswordResPage />
                </ProtectedAuth>
              }
            />
            <Route
              path="/"
              element={
                <Protected>
                  <SidePanel />
                  <UserDashboardPage />
                </Protected>
              }
            />
            <Route
              path="/UserDashbord"
              element={
                <Protected>
                  <SidePanel />
                  <UserDashboardPage />
                </Protected>
              }
            />
            <Route
              path="/ProfilePage"
              element={
                <Protected>
                  <SidePanel />
                  <ProfilePage />
                </Protected>
              }
            />
            <Route
              path="/UserList"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <UserList />
                </ProtectedAdmin>
              }
            />
            <Route
              path="/UserEdit/:UserId"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <UserEdit />
                </ProtectedAdmin>
              }
            />
            <Route
              path="/VideoList"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <VideoList />
                </ProtectedAdmin>
              }
            />
            <Route
              path="/VideoEdit/:VideoId"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <VideoEdit />
                </ProtectedAdmin>
              }
            />{' '}
            <Route
              path="/CategoriAdd"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <CategoriAdd />
                </ProtectedAdmin>
              }
            />
            <Route
              path="/AdminPanel"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <AdminPanel />
                </ProtectedAdmin>
              }
            />
            <Route
              path="/VideoAdd"
              element={
                <ProtectedAdmin>
                  <SidePanel />
                  <VideoAdd />
                </ProtectedAdmin>
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
