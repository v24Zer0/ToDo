import React from "react";
import { Pressable, Text, View } from "react-native";
import List from "../models/list";

interface ListProps {
    list: List;
    setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
    setModalList: React.Dispatch<React.SetStateAction<List>>;
    navigate(list: List): void;
}

// Add Pressable component to trigger modal display
// onPress triggers ItemList
// onLongPress triggers Modal display
const ListComponent: React.FC<ListProps> = ({ list, setModalVisible, setModalList, navigate }) => {
    return (
        <View>
            <Pressable onLongPress={() => {
                setModalVisible(true);
                setModalList(list);
            }} onPress={() => navigate(list)}>
                <Text>
                    {list.name}
                </Text>
            </Pressable>
        </View>
    );
}

export default ListComponent