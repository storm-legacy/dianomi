import React, { useState } from 'react';
import { Link, NavLink } from 'react-router-dom';
import { FaTh, FaUserAlt, FaBars } from 'react-icons/fa';
import { RiAdminFill } from 'react-icons/ri';
import { useAuthHelper } from '../../helpers/authHelper';
import { authAtom } from '../../states/auth';
import { useRecoilValue } from 'recoil';
import { BiLogOut } from 'react-icons/bi';
import '../../App.css';
const SidePanel = () => {
  const authHelper = useAuthHelper();
  const auth = useRecoilValue(authAtom);
  const isLogged = auth ? true : false;
  const [isOpen, setIsOpen] = useState(false);
  const toggle = () => setIsOpen(!isOpen);
  const menuItem = [
    { path: '/UserDashbord', name: 'Strona główna', icon: <FaTh /> },
    { path: '/ProfilePage', name: 'Progil', icon: <FaUserAlt /> },
    { path: '/AdminPanel', name: 'Admin', icon: <RiAdminFill /> },
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
          <button
            style={{
              backgroundColor: '#0D6EFD',
              width: '100%',
              height: '70%',
              color: 'white',
            }}
            className="link"
            onClick={() => {
              authHelper.logout({
                callback: () => {
                  window.location.reload();
                },
              });
            }}
          >
            <i className="icon">
              <BiLogOut></BiLogOut>
            </i>
            <p style={{ display: isOpen ? 'block' : 'none' }} className="link_text">
              wylogój
            </p>
          </button>
        </div>
      </div>
    </>
  );
};

export default SidePanel;
