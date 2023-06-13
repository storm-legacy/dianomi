import React, { useEffect, useState } from 'react';
import { AuthContext } from './context/AuthContext';
import { User } from './types/user.type';

import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import authService from './services/auth.service';

import ProtectedAdmin from './components/ProtectedAdmin';
import LoginPage from './pages/LoginPage/LoginPage';
import RegisterPage from './pages/RegisterPage/RegisterPage';
import NotFoundPage from './pages/NotFound/NotFound';

import UserDashboardPage from './pages/UserDashboard/UserDashboard';

import { ProfilePage } from './pages/ProfilePage/ProfilePage';
import SidePanel from './pages/SidePanel/SidePanel';
import { AdminPanel } from './pages/AdminPanel/AdminPanel';
import { VideoAdd } from './pages/AdminPanel/AdminPages/VideoAdd';
import UserList from './pages/AdminPanel/AdminPages/UserList';
import { VideoList } from './pages/AdminPanel/AdminPages/VideoList';
import { VideoEdit } from './pages/AdminPanel/AdminPages/VideoEdit';

import './App.css';
import { CanceledError } from 'axios';
import CategoriAdd from './pages/AdminPanel/AdminPages/CategoriAdd';
import { PasswordResPage } from './pages/LoginPage/PasswordResPage';
import { VideoPlayer } from './pages/VideoPlayerPage/VideoPlayer';
import { VideoComment } from './pages/VideoPlayerPage/CommentPage/VideoComment';

import { ResetPage } from './pages/ResetPage/ResetPage';
import { Report } from 'notiflix';
import PaymentPage from './pages/PaymentPage/PaymentPage';
import { ReportPage } from './pages/ReportPage/Report';
import { UserHistory } from './pages/ProfilePage/UserHistory';
import { CommentsList } from './pages/AdminPanel/AdminPages/CommentsList';

function App() {
  const [user, setUser] = useState<User | null>(JSON.parse(String(localStorage.getItem('user'))));
  // Information that user is not verified

  useEffect(() => {
    const { request, cancel } = authService.connectionCheck();

    if (user?.verified === false) {
      Report.info('E-mail verification', 'Please verify your email address, before you can proceed.', 'Okay');
      setUser(null);
    }

    if (user?.banned === true) {
      Report.warning('Account suspended', 'Your account was blocked. Please contact support in case it was a mistake.', 'I understand');
      setUser(null);
    }

    request
      .then(() => {
        console.info('Authorized');
      })
      .catch((err) => {
        if (err instanceof CanceledError) return;
        console.warn('Unauthorized');
      });

    return () => cancel();
  }, [user]);

  return (
    <>
      <AuthContext.Provider value={{ user, setUser }}>
        {user?.authToken && user?.verified ? (
          <div className="container-fluid d-flex m-0 p-0">
            <BrowserRouter>
              <SidePanel />
              <div className="container-fluid" style={{ maxHeight: '100vh' }}>
                <Routes>
                  <Route path="/" element={<UserDashboardPage />} />
                  <Route path="/UserDashbord" element={<UserDashboardPage />} />
                  <Route
                    path="/VideoPlayer/:VideoId"
                    element={
                      <>
                        <VideoPlayer />
                        <VideoComment />
                      </>
                    }
                  />
                  <Route path="/ProfilePage" element={<ProfilePage />} />
                  <Route path="/Report" element={<ReportPage />} />
                  <Route path="/History" element={<UserHistory />} />
                  <Route path="/payment" element={<PaymentPage />} />
                  <Route
                    path="/UserList"
                    element={
                      <ProtectedAdmin>
                        <UserList />
                      </ProtectedAdmin>
                    }
                  />
                  <Route
                    path="/VideoList"
                    element={
                      <ProtectedAdmin>
                        <VideoList />
                      </ProtectedAdmin>
                    }
                  />
                  <Route
                    path="/VideoEdit/:VideoId"
                    element={
                      <ProtectedAdmin>
                        <VideoEdit />
                      </ProtectedAdmin>
                    }
                  />
                  <Route
                    path="/CategoriAdd"
                    element={
                      <ProtectedAdmin>
                        <CategoriAdd />
                      </ProtectedAdmin>
                    }
                  />
                  <Route
                    path="/AdminPanel"
                    element={
                      <ProtectedAdmin>
                        <AdminPanel />
                      </ProtectedAdmin>
                    }
                  />
                  <Route
                    path="/VideoAdd"
                    element={
                      <ProtectedAdmin>
                        <VideoAdd />
                      </ProtectedAdmin>
                    }
                  />
                  <Route
                    path="/Comments"
                    element={
                      <ProtectedAdmin>
                        <CommentsList />
                      </ProtectedAdmin>
                    }
                  />
                  <Route path="*" element={<NotFoundPage />} />
                </Routes>
              </div>
            </BrowserRouter>
          </div>
        ) : (
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<Navigate replace to="/login" />} />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/register" element={<RegisterPage />} />
              <Route path="/reset-password" element={<ResetPage />} />
              <Route path="/password/reset" element={<PasswordResPage />} />
              <Route path="*" element={<Navigate replace to="/login" />} />
            </Routes>
          </BrowserRouter>
        )}
      </AuthContext.Provider>
    </>
  );
}

export default App;
