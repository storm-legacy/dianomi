import React, { useEffect, useState } from 'react';
import Modal from 'react-modal';
import adminService, { UserEditData, packages } from '../../../services/admin.service';

Modal.setAppElement('#root');

interface UserEditProps {
  isOpen: boolean;
  onRequestClose: () => void;
  userId: number | null;
  verified: boolean;
  OldEmail: string;
  packages: packages | undefined;
}

const UserEdit: React.FC<UserEditProps> = ({ isOpen, onRequestClose, userId, verified, OldEmail, packages }) => {
  const [userEmail, setUserEmail] = useState<string>('');
  const [isPasswordReset, setIsPasswordReset] = useState<boolean>(false);
  const [isVerified, setIsVerified] = useState(true);
  const [deletePack, setDeletePack] = useState<boolean>(false);
  const [packetId, setPacketId] = useState(packages?.id);
  const [isDisable, setIsDisabled] = useState<boolean>(false);
  const [tier, setTier] = useState<string | undefined>(packages?.tier);
  const [validFrom, setValidFrom] = useState<string>(String(packages?.valid_from));
  const [validUntil, setValidUntil] = useState<string>(String(packages?.valid_until));
  const customStyles = {
    content: {
      top: '50%',
      left: '50%',
      right: 'auto',
      bottom: 'auto',
      marginRight: '-50%',
      height: '95dvh',
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
    const packagesData: packages = {
      delete: deletePack,
      id: packetId,
      tier: tier,
      user_id: userId,
      valid_from: validFrom,
      valid_until: validUntil,
    };
    const DataUser: UserEditData = {
      email: userEmail,
      verified: isVerified,
      reset_password: isPasswordReset,
      packages: packagesData,
    };
    const { request } = adminService.patchUser(userId, DataUser);
    request
      .then((res) => {
        console.log(res);
        onRequestClose;
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
        <label>
          <input className="form-check-input" type="checkbox" name="packet" /> You want change packet
        </label>
        <br />
        <label>
          <input
            className="form-check-input"
            type="checkbox"
            name="delete"
            onChange={(event) => setDeletePack(event.target.checked)}
          />{' '}
          Delete
        </label>
        <br />
        <label>
          <p>ID:</p>
          <input
            className="form-control"
            value={packetId}
            onChange={(event) => setPacketId(parseInt(event.target.value))}
            type="number"
            name="id"
          />
        </label>
        <br />
        <label>
          <p> Tier:</p>
          <input className="form-control" type="text" name="tier" value={tier} />
        </label>
        <br />
        <br />
        <label>
          <p>Valid From:</p>
          <input className="form-control" type="data" name="valid_from" />
        </label>
        <br />
        <label>
          <p>Valid Until:</p>
          <input className="form-control" type="data" name="valid_until" />
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
