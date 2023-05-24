import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import adminService from '../../../services/admin.service';
import UserEdit from './UserEdit';
function UserList() {
  interface packages {
    delete: boolean;
    id: number;
    tier: string;
    user_id: number;
    valid_from: string;
    valid_until: string;
  }
  interface UserDataList {
    id: number;
    email: string;
    verified: boolean;
    packages: packages[];
  }
  const [userData, setUserData] = useState<UserDataList[]>([]);
  const [modalIsOpen, setModalIsOpen] = useState(false);
  const [userIdNow, setUserIdNow] = useState<number | null>();
  const [verified, setVerified] = useState<boolean>();
  const [userEmail, setUserEmail] = useState('');
  const [packagesUser, setPackagesUser] = useState<packages>();
  const openModal = (id: number, verified: boolean, email: string, packages: packages | undefined) => {
    setUserIdNow(id);
    setVerified(verified);
    setUserEmail(email);
    setPackagesUser(packages);
    setModalIsOpen(true);
  };

  const closeModal = () => {
    setModalIsOpen(false);
  };
  const DeleteUser = (value: number) => {
    console.log(value);
    const { request } = adminService.deleteUser(value);
    request.then((res) => console.log(res)).catch((err) => console.error(err.message));
    window.location.reload();
  };
  useEffect(() => {
    const { request } = adminService.takeUser();
    request
      .then((res) => {
        const data = res.data.map((data: UserDataList) => {
          const updatedPackages = data.packages.map((pkg: packages) => {
            const { delete: pkgDelete, id: pkgId, tier, user_id, valid_from, valid_until } = pkg;
            return { delete: pkgDelete, id: pkgId, tier, user_id, valid_from, valid_until };
          });
          console.log(updatedPackages);
          return { id: data.id, email: data.email, verified: data.verified, packages: updatedPackages };
        });
        setUserData(data);
        console.log(userData);
      })
      .catch((err) => {
        console.error(err.message);
      });
  }, []);
  return (
    <>
      <div className="text-center MyList">
        <h1>User List</h1>
        <div className="row myRow myRowCont overflow-auto">
          <div className="row myRow">
            <div className="col-1 border border-primary">ID</div>
            <div className="col-2 border border-primary">Email</div>
            <div className="col-2 border border-primary">role</div>

            <div className="col-2  border border-primary">Valid From</div>
            <div className="col-2  border border-primary">Valid until</div>
            <div className="col-1  border border-primary">Verified</div>
            <div className="col-1 border border-primary"> Settings</div>
          </div>
          {userData.map((item, index) => (
            <div className="row" key={index}>
              <div className="col-1 border border-primary">{item.id}</div>
              <div className="col-2 border border-primary">{item.email}</div>
              {item.packages.length > 0 ? (
                item.packages.map((pkg, index) => (
                  <React.Fragment key={index}>
                    <div className="col-2 border border-primary">{pkg.tier}</div>
                    <div className="col-2 border border-primary">{pkg.valid_from}</div>
                    <div className="col-2 border border-primary">{pkg.valid_until}</div>
                  </React.Fragment>
                ))
              ) : (
                <>
                  <div className="col-2 border border-primary">empty</div>
                  <div className="col-2 border border-primary">empty</div>
                  <div className="col-2 border border-primary">empty</div>
                </>
              )}
              <div className="col-1 border border-primary">{item.verified ? 'Yes' : 'No'}</div>
              <div className="col-1 border border-primary">
                {' '}
                <button className="custom-link link-primary" onClick={() => DeleteUser(item.id)}>
                  {' '}
                  Delete{' '}
                </button>{' '}
                <button
                  onClick={() => openModal(item.id, item.verified, item.email, item.packages[0])}
                  className="custom-link link-primary"
                >
                  Edit
                </button>{' '}
                Verified
              </div>
            </div>
          ))}
          <UserEdit
            isOpen={modalIsOpen}
            onRequestClose={closeModal}
            userId={userIdNow!}
            verified={verified!}
            OldEmail={userEmail}
            packages={packagesUser}
          />
        </div>
      </div>
    </>
  );
}

export default UserList;
