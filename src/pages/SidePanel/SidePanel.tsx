import React, { useContext, useState } from 'react';
import { NavLink, useNavigate } from 'react-router-dom';
import { FaTh, FaUserAlt, FaBars } from 'react-icons/fa';
import { RiAdminFill } from 'react-icons/ri';
import { BiLogOut } from 'react-icons/bi';
import '../../App.css';
import authService from '../../services/auth.service';
import { AuthContext } from '../../context/AuthContext';

const SidePanel = () => {
  const navigate = useNavigate();
  const [isOpen, setIsOpen] = useState(false);
  const toggle = () => setIsOpen(!isOpen);

  const { user } = useContext(AuthContext);

  let menuItem = [
    { path: '/UserDashbord', name: 'Home', icon: <FaTh /> },
    { path: '/ProfilePage', name: 'Profile', icon: <FaUserAlt /> },
  ];

  if (user?.role === 'administrator') {
    menuItem = [...menuItem, { path: '/AdminPanel', name: 'Admin', icon: <RiAdminFill /> }];
  }

  const handleLogout = () => {
    const { request } = authService.logout();
    request
      .then(() => {
        console.info('Logout');
        localStorage.clear();
        location.reload();
      })
      .catch((err) => {
        console.error(err.message);
        localStorage.clear();
        location.reload();
      });
  };

  return (
    <>
      <div style={{ width: isOpen ? '200px' : '50px' }} className="sidebar">
        <div className="top_section">
          <h1 style={{ display: isOpen ? 'block' : 'none' }} className="logo">
            Logo
          </h1>
          <div style={{ marginLeft: isOpen ? '50px' : '0px' }} className="bars">
            <FaBars onClick={toggle} />
          </div>
        </div>
        {menuItem.map((item, index) => (
          <NavLink to={item.path} key={index} className="link">
            <div className="icon">{item.icon}</div>
            <div style={{ display: isOpen ? 'block' : 'none' }} className="link_text">
              {item.name}
            </div>
          </NavLink>
        ))}
        <button
          style={{
            backgroundColor: '#0D6EFD',
            width: '100%',
            height: '70%',
            color: 'white',
          }}
          className="link"
          onClick={handleLogout}
        >
          <i className="icon">
            <BiLogOut></BiLogOut>
          </i>
          <p style={{ display: isOpen ? 'block' : 'none' }} className="link_text">
            Logout
          </p>
        </button>
      </div>
    </>
  );
};

export default SidePanel;
