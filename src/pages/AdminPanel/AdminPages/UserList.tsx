import React from 'react';
import { Link } from 'react-router-dom';
function UserList() {
  const UserListItem = [
    {
      id: '1',
      email: 'kick1332@gmail.com',
      password: 'M@rcel1332',
      verified_at: '12-12-2000',
      updated_at: '12-12-2000',
      created_at: '12-12-2000',
    },
    {
      id: '2',
      email: 'Havajska13@gmail.com',
      password: 'Pomidorowa',
      verified_at: '11-11-2000',
      updated_at: '11-11-2000',
      created_at: '11-11-2000',
    },
    {
      id: '3',
      email: 'Wokulski@gmail.com',
      password: 'Kochamłęcką',
      verified_at: '11-01-2000',
      updated_at: '11-01-2000',
      created_at: '11-01-2000',
    },
  ];
  function DealeatUser() {
    console.log('');
  }
  return (
    <>
      <div className="text-center">
        <div className=" position-absolute top-50 start-50 translate-middle">
          <div className="row myRow myRowCont overflow-auto">
            <h1>User List</h1>
            <div className="row myRow">
              <div className="col-1 border border-primary">ID</div>
              <div className="col-3 border border-primary">Email</div>
              <div className="col-2  border border-primary">Created at</div>
              <div className="col-2  border border-primary">Updated at</div>
              <div className="col-2  border border-primary">Verified at</div>
              <div className="col-2 border border-primary"> Settings</div>
            </div>
            {UserListItem.map((item, index) => (
              <div className="row" key={index}>
                <div className="col-1 border border-primary">{item.id}</div>
                <div className="col-3 border border-primary">{item.email}</div>
                <div className="col-2 border border-primary">{item.created_at}</div>
                <div className="col-2 border border-primary">{item.updated_at}</div>
                <div className="col-2 border border-primary">{item.verified_at}</div>
                <div className="col-2 border border-primary">
                  {' '}
                  Delete <Link to={'/UserEdit/' + item.id}>Edit</Link> Ban
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}

export default UserList;
