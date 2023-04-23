import React, { useState } from 'react';
import { Link, NavLink } from 'react-router-dom';
import { FaTh, FaUserAlt, FaBars } from 'react-icons/fa';
import { useAuthHelper } from '../../helpers/authHelper';
import { authAtom } from '../../states/auth';
import { useRecoilValue } from 'recoil';

const SidePanel = () => {
  const authHelper = useAuthHelper();
  const auth = useRecoilValue(authAtom);
  const isLogged = auth ? true : false;
  const [isOpen, setIsOpen] = useState(false);
  const toggle = () => setIsOpen(!isOpen);
  const menuItem = [
    { path: '/UserDashbord', name: 'dashbord', icon: <FaTh /> },
    { path: '/ProfilePage', name: 'ProfilePage', icon: <FaUserAlt /> },
  ];
  return (
    <>
      <div className="container">
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
        </div>
      </div>
      <button
        className="position-absolute top-0 start-100 translate-middle"
        onClick={() => {
          authHelper.logout({
            callback: () => {
              window.location.reload();
            },
          });
        }}
      >
        Logout
      </button>
    </>
  );
};

export default SidePanel;
