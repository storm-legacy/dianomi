import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

const ProtectedAuth = ({ children }: { children: ReactNode }) => {
  const token = localStorage.getItem('token');
  if (token) {
    return <Navigate to="/" replace />;
  }

  return <>{children}</>;
};

export default ProtectedAuth;
