import React, { useEffect, useState } from 'react';
import Modal from 'react-modal';
import adminService, { UserEditData } from '../../../services/admin.service';

Modal.setAppElement('#root');

interface UserEditProps {
  isOpen: boolean;
  onRequestClose: () => void;
  userId: number | null;
  verified: boolean;
  OldEmail: string;
}

const UserEdit: React.FC<UserEditProps> = ({ isOpen, onRequestClose, userId, verified, OldEmail }) => {
  const [userEmail, setUserEmail] = useState<string>('');
  const [isPasswordReset, setIsPasswordReset] = useState<boolean>(false);
  const [isVerified, setIsVerified] = useState(true);
  const [isDisable, setIsDisabled] = useState<boolean>(false);
  const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      height: '50dvh',
      width: '50dvh',
      transform: 'translate(-50%, -50%)',
    },
  };
  useEffect(() => {
    if (!/^\S+@\S+\.\S+$/.test(userEmail)) {
      setIsDisabled(true);
    } else {
      setIsDisabled(false);
    }
  }, [userEmail]);

  const handleSubmit = (event: any) => {
    event.preventDefault();
    const DataUser: UserEditData = {
      email: userEmail,
      verified: isVerified,
      reset_password: isPasswordReset,
      packages: [],
    };
    const { request } = adminService.patchUser(userId, DataUser);
    request
      .then((res) => {
        console.log(res);
        window.location.reload();
      })
      .catch((err) => {
        console.error(err.message);
      });
  };
  return (
    <Modal isOpen={isOpen} onRequestClose={onRequestClose} contentLabel="Modal" style={customStyles}>
      <form style={{ margin: '25px' }} onSubmit={handleSubmit}>
        {' '}
        <h2>User edit</h2>
        <h3>Email</h3>
        <label>
          <input
            className="form-control"
            type="text"
            value={userEmail}
            placeholder={OldEmail}
            onChange={(event) => setUserEmail(event.target.value)}
          />
        </label>
        <br />
        <label className="form-check-label">
          <input
            className="form-check-input"
            type="checkbox"
            value=""
            id="Very"
            checked={isVerified || false}
            onChange={(e) => setIsVerified(e.target.checked)}
          />
          Verified
        </label>
        <br />
        <label className="form-check-label">
          <input
            className="form-check-input"
            type="checkbox"
            value=""
            id="ResPas"
            onChange={(e) => setIsPasswordReset(e.target.checked)}
          />
          Password resetset
        </label>
        <br />
        <button className="btn btn-primary" disabled={isDisable}>
          Edit
        </button>
      </form>

      <button style={{ margin: '25px' }} className="btn btn-primary" onClick={onRequestClose}>
        Close
      </button>
    </Modal>
  );
};

export default UserEdit;
