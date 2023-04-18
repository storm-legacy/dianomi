import React, { ReactNode } from 'react';
import { Navigate } from 'react-router-dom';

interface Props {
  isLoggedIn: boolean;
  children: ReactNode;
}

const Protected = ({ isLoggedIn, children }: Props) => {
  if (!isLoggedIn) return <Navigate to="/login" replace />;
  return <>{children}</>;
};

export default Protected;
