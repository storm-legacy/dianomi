import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

const ProtectedAdmin = ({ children }: { children: ReactNode }) => {
  // Check if logged in
  const token = localStorage.getItem('token');
  const role = localStorage.getItem('role');

  if (!token || !role) {
    return <Navigate to="/login" replace />;
  }

  // Check if administrator
  if (String(role) !== 'administrator') {
    return <Navigate to="/" replace />;
  }

  return <>{children}</>;
};

export default ProtectedAdmin;
