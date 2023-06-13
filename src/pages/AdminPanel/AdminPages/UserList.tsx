import React, { useEffect, useState, useContext } from 'react';
import adminService, { Package } from '../../../services/admin.service';
import UserEdit from './UserEdit';
import { Report, Notify, Confirm, Block } from 'notiflix';
import { FaTrash, FaWrench, FaUserTimes, FaUserCircle } from 'react-icons/fa';
import { AuthContext } from '../../../context/AuthContext';

interface UserDataList {
  id: number;
  email: string;
  verified: boolean;
  role: string;
  banned: boolean;
}

const UserList = () => {
  const user = useContext(AuthContext);
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

  const refreshList = () => {
    const { request } = adminService.takeUser();
    request
      .then(({ data }) => {
        setUserData(data);
      })
      .catch((err) => {
        Report.failure('Could not fetch users list', `Server returned message: ${err.message}`, 'Okay');
      });
  }

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
            refreshList();
          })
          .catch((err) => {
            Notify.failure('User could not be removed!');
          });
      },
      () => {
        return;
      },
      {
        titleColor: '#C1292E',
        messageColor: '#C1292E',
        okButtonBackground: '#C1292E'
      },
    );
  };

  const handleBanUser = (userId: number) => {
    Confirm.show(
      'Ban user',
      'This user\'s account will be suspended. Do you wish to proceed?',
      'Yes',
      'No',
      () => {
        const { request } = adminService.postBanUser(userId);
        request.then(() => {
          Notify.success("Account was successfully suspended!");
          refreshList();
        })
          .catch((err) => {
            Notify.failure('Error occured while account was about to be blocked!');
          });
      },
      () => {
        return;
      },
      {},
    );
  }

  const handleUnbanUser = (userId: number) => {
    Confirm.show(
      'Unban user',
      'This user\'s account will be recovered from suspension. Do you wish to proceed?',
      'Yes',
      'No',
      () => {
        const { request } = adminService.postUnbanUser(userId);
        request.then(() => {
          Notify.success("Account was successfully resumed!");
          refreshList();
        })
          .catch((err) => {
            Notify.failure('Error occured while account was about to be resumed!');
          });
      },
      () => {
        return;
      },
      {},
    );
  }

  useEffect(() => {
    refreshList();
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

                  <button className="btn btn-secondary mx-2" onClick={() => openModal(item.id)}>
                    <FaWrench />
                  </button>
                  {item.email !== user.user?.email ? (
                    <>
                      <button className="btn btn-warning mx-2" onClick={!item.banned ? () => handleBanUser(item.id) : () => handleUnbanUser(item.id)}>
                        {!item.banned ? (<FaUserTimes />) : (<FaUserCircle />)}
                      </button>
                      <button className="btn btn-danger mx-2" onClick={() => DeleteUser(item.id)}>
                        <FaTrash />
                      </button>
                    </>
                  ) : (
                    <>
                      <button className="btn btn-warning mx-2" disabled>
                        <FaUserCircle />
                      </button>
                      <button className="btn btn-danger mx-2" disabled>
                        <FaTrash />
                      </button>
                    </>
                  )}
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
