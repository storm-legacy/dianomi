import React from 'react';
import { NavLink } from 'react-router-dom';
export const AdminPanel = () => {
  return (
    <div>
      <NavLink to={'/VideoAdd'}> dodawanie wideo</NavLink>
    </div>
  );
};
