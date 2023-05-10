import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

const Protected = ({ children }: { children: ReactNode }) => {
  // Check if logged in
  const token = localStorage.getItem('token');
  const role = localStorage.getItem('role');

  if (!token || !['administrator', 'premium', 'free'].includes(String(role))) {
    return <Navigate to="/login" replace />;
  }

  return <>{children}</>;
};

export default Protected;
