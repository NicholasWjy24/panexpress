import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:panexpress/data/model/menu_item.dart';
import 'package:panexpress/data/repository/menu_repository.dart';
import 'package:panexpress/data/service/menu_service.dart';

sealed class HomeEvent {
  const HomeEvent();
}

class HomeStarted extends HomeEvent {
  const HomeStarted();
}

class HomeMenuRefreshRequested extends HomeEvent {
  const HomeMenuRefreshRequested();
}

sealed class HomeState extends Equatable {
  const HomeState();

  @override
  List<Object?> get props => [];
}

class HomeInitial extends HomeState {
  const HomeInitial();
}

class HomeLoading extends HomeState {
  const HomeLoading();
}

class HomeSuccess extends HomeState {
  const HomeSuccess(this.menus);

  final List<MenuItem> menus;

  @override
  List<Object?> get props => [menus];
}

class HomeFailure extends HomeState {
  const HomeFailure(this.message);

  final String message;

  @override
  List<Object?> get props => [message];
}

class HomeBloc extends Bloc<HomeEvent, HomeState> {
  HomeBloc(this._menuRepository) : super(const HomeInitial()) {
    on<HomeStarted>(_loadMenus);
    on<HomeMenuRefreshRequested>(_loadMenus);
  }

  final MenuRepository _menuRepository;

  Future<void> _loadMenus(
    HomeEvent event,
    Emitter<HomeState> emit,
  ) async {
    emit(const HomeLoading());

    try {
      final menus = await _menuRepository.getMenus();
      emit(HomeSuccess(menus));
    } on MenuException catch (error) {
      emit(HomeFailure(error.message));
    } catch (_) {
      emit(const HomeFailure('Could not load menu data.'));
    }
  }
}
