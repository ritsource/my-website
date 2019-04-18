import React, { useState } from 'react';
import { Link, NavLink } from 'react-router-dom';
import { IoIosSunny, IoIosMoon } from 'react-icons/io';

import toggleTheme from '../utils/changeTheme';

const Header = () => {
	const [ isLight, setIsLight ] = useState(true);

	return (
		<div className="Header-c-00 Flex-Row-Space-Between">
			<Link to="/">
				<h3>
					<span>🐳</span>
				</h3>
			</Link>

			<div className="Header-Navbtn-Container-Div-01 Flex-Row-Center">
				{!isLight ? (
					<button onClick={() => toggleTheme(isLight, setIsLight)} className="Flex-Row-Center">
						<IoIosSunny style={{ color: 'var(--text-color)', fontSize: '24px' }} />
					</button>
				) : (
					<button onClick={() => toggleTheme(isLight, setIsLight)} className="Flex-Row-Center">
						<IoIosMoon style={{ color: 'var(--text-color)', fontSize: '24px' }} />
					</button>
				)}
			</div>
		</div>
	);
};

export default Header;
