import React, { useEffect, useState } from 'react';
import adminService, { Package } from '../../../services/admin.service';
import UserEdit from './UserEdit';
import { Report, Notify, Confirm, Block } from 'notiflix';
import { FaTrash, FaWrench } from 'react-icons/fa';

interface UserDataList {
  id: number;
  email: string;
  verified: boolean;
  role: string;
}

const UserList = () => {
  const [userData, setUserData] = useState<UserDataList[]>([]);
  const [modalIsOpen, setModalIsOpen] = useState(false);
  const [userIdNow, setUserIdNow] = useState<number>(0);

  const openModal = (userId: number) => {
    setUserIdNow(userId);
    setModalIsOpen(true);
  };

  const closeModal = () => {
    setModalIsOpen(false);
  };

  const DeleteUser = (userId: number) => {
    Confirm.show(
      'User will be removed!',
      'Are you sure you would like to proceed?',
      'Yes',
      'No',
      () => {
        const { request } = adminService.deleteUser(userId);
        request
          .then((res) => {
            Notify.success('User was successfuly removed!');
            window.location.reload();
          })
          .catch((err) => {
            Notify.failure('User could not be removed!');
          });
      },
      () => {
        return;
      },
      {},
    );
  };

  useEffect(() => {
    const { request } = adminService.takeUser();
    request
      .then(({ data }) => {
        setUserData(data);
      })
      .catch((err) => {
        Report.failure('Could not fetch users list', `Server returned message: ${err.message}`, 'Okay');
      });
  }, []);

  return (
    <>
      <div className="text-center MyList">
        <h1>User List</h1>
        <div className="row myRow myRowCont overflow-auto">
          <div className="row myRow">
            <div className="col-1 border border-primary">ID</div>
            <div className="col-3 border border-primary">Email</div>
            <div className="col-3 border border-primary">Role</div>
            <div className="col-2  border border-primary">Verified</div>
            <div className="col-2 border border-primary"> Settings</div>
          </div>
          {userData.map((item, index) => {
            return (
              <div className="row" key={index}>
                <div className="col-1 p-2 border border-primary">{item.id}</div>
                <div className="col-3 p-2 border border-primary">{item.email}</div>
                <div className="col-3 p-2 border border-primary">{item.role}</div>
                <div className="col-2 p-2 border border-primary">{item.verified ? 'Yes' : 'No'}</div>
                <div className="col-2 p-2 border border-primary d-flex flex-row">
                  <button className="btn btn-warning mx-2" onClick={() => openModal(item.id)}>
                    <FaWrench />
                  </button>
                  <button className="btn btn-danger mx-2" onClick={() => DeleteUser(item.id)}>
                    <FaTrash />
                  </button>
                </div>
              </div>
            );
          })}
          <UserEdit isOpen={modalIsOpen} onRequestClose={closeModal} userId={userIdNow} />
        </div>
      </div>
    </>
  );
};

export default UserList;
