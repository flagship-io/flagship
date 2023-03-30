import React from "react";

export const MyReactComponent = () => {
    const backgroundColorFlag = instance.getFeatureVariables('OPT-flag-jsx')

    return (
        <div
            style={{
                height: "200px",
                width: "200px",
                backgroundColor: backgroundColorFlag.getValue(),
            }}
        >
            {"I'm a square with color=" + backgroundColorFlag.getValue()}
            <button style={{color: btnColorFlag.getValue()}}>Button</button>
        </div>
    );
};
