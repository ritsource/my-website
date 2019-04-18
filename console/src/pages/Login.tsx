import React from 'react';
import { FaGoogle } from 'react-icons/fa';

const LoginPage = () => {
	return (
		<div className="Page-c-00 Flex-Column-Center">
			<div className="Page-Container-00 Flex-Column-Center">
				<a href={process.env.REACT_APP_API_URL + '/api/auth/google'}>
					<button
						style={{
							fontSize: '14px',
							letterSpacing: '0.1px',
							marginTop: '-70px',
							padding: '10px 20px',
							fontWeight: 'bold'
						}}
						className="Theme-Btn-Main Flex-Row-Center"
					>
						<FaGoogle style={{ marginRight: '10px' }} />Login as Ritwik
					</button>
				</a>
			</div>
		</div>
	);
};

export default LoginPage;
