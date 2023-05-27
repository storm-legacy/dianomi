import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';
import { User } from '../types/user.type';

const ProtectedAdmin = ({ children }: { children: ReactNode }) => {
  // Check if logged in
  const user: User = JSON.parse(String(localStorage.getItem('user')));

  if (!user.authToken || !user.role) {
    return <Navigate to="/login" replace />;
  }

  // Check if administrator
  if (user.role !== 'administrator') {
    return <Navigate to="/" replace />;
  }

  return <>{children}</>;
};

export default ProtectedAdmin;
