import * as React from 'react';

// tslint:disable-next-line:interface-name
interface CheckBoxProps {
    checked:boolean;
    msg:string;
    checkBoxClick:(checked:boolean)=>void;
}

export default class CheckBox extends React.Component<CheckBoxProps> {

    // 아이콘 사이트
    // expo vector icons expo.github.io/vector-icons/

    public render() {
        const handleCheck = this.handleCheck.bind(this);
        return (
            <div>
                <input type="checkbox" onChange={handleCheck} defaultChecked={this.props.checked}/>
                { !!this.props.msg ? this.props.msg : '' }
            </div>
        );
    }

    private handleCheck() {
        this.props.checkBoxClick(!this.props.checked);
    }
}