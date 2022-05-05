import React from "react";
import { Text, View } from "react-native";
import List from "../models/list";

interface ListProps {
    list: List;
}

const ListComponent: React.FC<ListProps> = ({ list }) => {
    return (
        <View>
            <Text>
                {list.name}
            </Text>
        </View>
    );
}

export default ListComponent