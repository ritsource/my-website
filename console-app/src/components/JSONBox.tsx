import React, { useState, useEffect } from 'react';
import { IoMdCheckmarkCircleOutline } from 'react-icons/io';

import { Data } from '../data/data';

type MyProps = {
	object: Data;
};

const JSONBox = (props: MyProps) => {
	const { object } = props;

	const [ json, setJson ] = useState(JSON.stringify(object, null, 4));
	const [ isValid, setIsValid ] = useState(false);
	const [ isAsync, setIsAsync ] = useState(false);

	const [ jsonBoxEd, setJsonBoxEd ] = useState(false);

	useEffect(
		() => {
			setJson(JSON.stringify(object, null, 4));
			const valid = validateJson(json);
			setIsValid(valid);
		},
		[ object ]
	);

	const validateJson = (str: string) => {
		try {
			JSON.parse(str);
		} catch (e) {
			// console.log(e.message.match(/\d+/g));
			return false;
		}
		return true;
	};

	return (
		<div className="JOSNBox-c-00 SearchBox-c-00 Theme-Box-Shadow">
			<h4 className="Flex-Row-Space-Between">
				JSON Data {isAsync && <div className="Theme-Loading-Spin-Div" />}
			</h4>
			<div className="Flex-Row-Space-Between">
				<p className="SearchBox-Description-P-01">Quickly modify using JSON data structure.</p>
				{/* {isValid ? 'True' : 'False'} */}
				<IoMdCheckmarkCircleOutline
					style={
						isValid ? (
							{ color: 'var(--safe-green-color)', fontSize: '20px', marginRight: '20px' }
						) : (
							{ color: 'var(--danger-red-color)', fontSize: '20px', marginRight: '20px' }
						)
					}
				/>
			</div>
			<textarea
				value={json}
				// style={jsonBoxEd ? {} : { border: '1px solid var(--border-color)' }}
				onChange={(e) => {
					if (jsonBoxEd) {
						const valid = validateJson(e.target.value);
						setIsValid(valid);
						setJson(e.target.value);
					}
				}}
			/>

			<div
				style={{
					width: 'calc(100% - 40px)',
					padding: '10px 20px 20px 20px'
				}}
				className="Flex-Row-Start"
			>
				{jsonBoxEd ? (
					<React.Fragment>
						<button className="Theme-Btn-Green">Save</button>
						<button
							style={{ marginLeft: '12px' }}
							className="Theme-Btn-Grey"
							onClick={() => {
								setJsonBoxEd(false);
								setJson(JSON.stringify(object, null, 4));
							}}
						>
							Cancel
						</button>
					</React.Fragment>
				) : (
					<button className="Theme-Btn-Main" onClick={() => setJsonBoxEd(true)}>
						Edit
					</button>
				)}
			</div>
		</div>
	);
};

export default JSONBox;
